package serialize

/**
 *	序列化辅助工具
 *	传输的数据格式为：[头部][数据内容]，其中头部信息为数据内容的字节长度，
 *
 *	写入时，先获取到数据内容的字节长度，再将长度按int32类型(固定4字节)写入缓存，再将数据内容写入缓存，
 *	构造好传输的字节数据后返回，用于socket传输
 *
 *	解析时，先读取4个字节(1个int32大小)，转换成int32为传输数据的字节大小，再根据数据的字节大小从socket获取
 *	到传输的字节数据，再通过proto转成对应的类型
 *
 */
import (
	"github.com/golang/protobuf/proto"
	"bytes"
	"encoding/binary"
	"net"
	"io"
)

func ToBytes(data proto.Message) ([]byte, error) {
	//proto处理数据
	pData, err := proto.Marshal(data)
	if err != nil {
		return []byte{}, err
	}

	buff := bytes.NewBuffer([]byte{})
	//获数据长度
	size := int32(len(pData))
	//将数据长度写入头部，以int32格式写入，占4个字节
	binary.Write(buff, binary.BigEndian, size)
	//写入proto数据
	buff.Write(pData)

	return buff.Bytes(), nil
}

func ToProto(conn net.Conn, pb proto.Message) error {
	//读取数据头，得到数据的字节长度
	size, err := readInt32(conn)
	if err != nil {
		return err
	}

	//以传输的数据长度生成字节数组
	buff := make([]byte, size)
	//从socket获取对应长度数据
	_, err = io.ReadFull(conn, buff)
	if err != nil {
		return err
	}
	//以proto反序列化数据
	return proto.Unmarshal(buff, pb)
}

func readInt32(conn net.Conn) (int32, error) {
	buff := make([]byte, 4, 4)
	_, err := conn.Read(buff)
	if err != nil {
		return -1, err
	}

	bb := bytes.NewBuffer(buff)

	var res int32

	binary.Read(bb, binary.BigEndian, &res)

	return res, nil
}
