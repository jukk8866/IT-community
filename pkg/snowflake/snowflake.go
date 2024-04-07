package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

// 在初始化 (Init) 函数中，它接收一个起始时间和机器ID作为参数。
// 其中起始时间是一个字符串形式的日期（"2006-01-02" 格式），
// 通过解析该起始时间并转换为 time.Time 类型来确定算法中的时间起点。机器ID是一个整数类型。
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	//首先声明了一个 time.Time 变量 st，然后使用 time.Parse 函数将起始时间字符串
	//解析为 time.Time 类型， 并将结果赋值给 st
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	//起始时间转换为纳秒级的 Unix 时间戳除以1000000，来设置雪花算法中的 Epoch （时间起点）。然后使用 sf.NewNode 函数，
	//根据机器ID创建一个雪花算法节点，并将结果赋值给全局的 node 变量。如果创建节点时出现错误，将返回错误信息。在生成（GenID）函数中，
	//它使用全局的 node 变量通过调用 Generate 方法来生成一个雪花算法生成的唯一ID，并将其转换为 int64 类型并返回。
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}
