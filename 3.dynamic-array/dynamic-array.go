package dynamicarray

func DynamicArray(n int32, queries [][]int32) []int32 {
    // Write your code here
    seqList := make([][]int32, n)
    for i := range seqList {
        seqList[i] = []int32{}
    }

    var lastAnswer int32
    results := []int32{}
    
    for _, query := range queries {
        queryType := query[0]
        x := query[1]
        y := query[2]
        seqIndex := (x ^ lastAnswer) % n
        if queryType == 1 {
            seqList[seqIndex] = append(seqList[seqIndex], y)
        }else if queryType == 2 {
            // Query tipe 2: ambil elemen dan perbarui lastAnswer
            element := seqList[seqIndex][y % int32(len(seqList[seqIndex]))]
            lastAnswer = element
            results = append(results, lastAnswer)
        }
    }
        return results
}