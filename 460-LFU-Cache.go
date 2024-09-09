type LFUCache struct {
    size int
    capacity int
    minFreq int
    freq map[int]*list.List
    eles map[int]*list.Element
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
\t\tcapacity: capacity,
\t\tfreq:     make(map[int]*list.List),
\t\teles:     make(map[int]*list.Element),
\t}
}

func (c *LFUCache) Get(key int) int {
    ele, exists := c.eles[key]
    if !exists {
        return -1
    }

    data := ele.Value.(Data)
    c.incrementFrequency(data, key)

    return data.Val
}

func (c *LFUCache) Put(key int, value int) {
    if c.capacity == 0 {
        return
    }

    if ele, exists := c.eles[key]; exists {
        data := ele.Value.(Data)
        data.Val = value
        c.incrementFrequency(data, key)
        return
    }

    if c.IsFull() {
        c.evict()
    }

    c.add(key, value)
}

func (c *LFUCache) incrementFrequency(data Data, key int) {
    li := c.freq[data.Freq]
    li.Remove(c.eles[key])

    if li.Len() == 0 && data.Freq == c.minFreq {
        c.minFreq++
    }

    data.Freq++
    c.addToFreqList(data.Freq, data, key)
}

func (c *LFUCache) IsFull() bool {
    return c.size == c.capacity
}

func (c *LFUCache) evict() {
\tli := c.freq[c.minFreq]
\tevicted := li.Remove(li.Front()).(Data)

\tdelete(c.eles, evicted.Key)
\tif li.Len() == 0 {
\t\tdelete(c.freq, c.minFreq)
\t}

\tc.size--
}

func (c *LFUCache) add(key , val int) {
    data := NewData(key, val)
    c.addToFreqList(1, data, key)
    c.minFreq = 1
    c.size++
}

func (c *LFUCache) addToFreqList(freq int, data Data, key int) {
    if _, exists := c.freq[freq]; !exists {
        c.freq[freq] = list.New()
    }
    c.eles[key] = c.freq[freq].PushBack(data)
}

type Data struct {
    Key int
    Val int
    Freq int
}

func NewData(key, val int) Data {
    return Data{
        Key: key,
        Val: val,
        Freq: 1,
    }
}