package algorithm

/**
	基于数组实现的循环队列
 */

type ArrayLoopQueue struct {
	data []interface{}
	head int
	tail int
	size int
}

func (alq *ArrayLoopQueue)Add(value interface{}) bool{
	if alq.data == nil{
		alq.data = make([]interface{},10,10)
	}
	length := len(alq.data)
	if (alq.tail+1) % length == alq.head{
		return false
	}
	alq.data[alq.tail % length] = value
	alq.tail++
	alq.size++
	return true
}

func (alq *ArrayLoopQueue) Take() interface{}{
	length := len(alq.data)
	if alq.tail == alq.head {
		return nil
	}
	v := alq.data[alq.head % length]
	alq.size--
	alq.head++
	return v
}
