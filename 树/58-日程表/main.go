package main

type MyCalendar struct {
	events []Event
}

type Event struct {
	start int
	end   int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start int, end int) bool {
	//找最后一个结束事件小于等于start的事件
	//idx := c.halfsearch(start, end,0, len(c.events)-1)
	idx := -1
	startIndex := 0
	endIndex := len(c.events) - 1
	///最后一个结束时间小于 要插入的开始时间
	for endIndex >= startIndex {
		mid := startIndex + (endIndex-startIndex)/2
		if c.events[mid].end <= start {
			if mid == len(c.events)-1 || c.events[mid+1].end > start {
				idx = mid
				break
			}
			startIndex = mid + 1
		} else {
			endIndex = mid - 1
		}
	}
	event := Event{start: start, end: end}
	//没找到的话，判断第一个数据是否大于要插入的值
	if idx == -1 {
		if len(c.events) == 0 {
			c.events = append(c.events, event)
			return true
		} else {
			firstEvent := c.events[0]
			if firstEvent.start >= end {
				temp := c.events
				c.events = append([]Event{}, event)
				c.events = append(c.events, temp...)
				return true
			}
		}
	} else {
		if idx == len(c.events)-1 {
			c.events = append(c.events, event)
			return true
		} else {
			nextEvent := c.events[idx+1]
			if nextEvent.start >= end {
				temp := c.events
				c.events = append([]Event{}, temp[:idx+1]...)
				c.events = append(c.events, event)
				c.events = append(c.events, temp[idx+1:]...)
				return true
			} else {
				return false
			}
		}
	}
	return false
}
