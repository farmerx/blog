# SORT
------

## 冒泡排序（bubble sort）

> 冒泡排序方法是最简单的排序方法。这种方法的基本思想是，将待排序的元素看作是竖着排列的“气泡”，较小的元素比较轻，从而要往上浮。在冒泡排序算法中我们要对这个“气泡”序列处理若干遍。所谓一遍处理，就是自底向上检查一遍这个序列，并时刻注意两个相邻的元素的顺序是否正确。如果发现两个相邻元素的顺序不对，即“轻”的元素在下面，就交换它们的位置。显然，处理一遍之后，“最轻”的元素就浮到了最高位置；处理二遍之后，“次轻”的元素就浮到了次高位置。在作第二遍处理时，由于最高位置上的元素已是“最轻”元素，所以不必检查。一般地，第i遍处理时，不必检查第i高位置以上的元素，因为经过前面i-1遍的处理，它们已正确地排好序。 
> 冒泡排序是稳定的。算法时间复杂度是O(n ^2) 

```
func  bubblesort(array []int) []int {
	length := len(array)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
```

## 选择排序

> 选择排序，就是假设最外层的循环第一个就是最小值的下标， 以后内存循环的时候，如果有值小于它，修改最小值的下标
```
func selectSort(arr []int){
  for i:=0;i<len(arr);i++{
      min:=i
      for j:=i+1;j<len(arr);j++{
         if arr[j]<arr[min]{
           m = j
         }
      }
      arr[i],arr[min] = arr[min],arr[j]
  }
}

```

## 插入排序
> 插入排序的基本思想是，经过i-1遍处理后,L[1..i-1]己排好序。第i遍处理仅将L[i]插入L[1..i-1]的适当位置，使得L[1..i]又是排好序的序列。要达到这个目的，我们可以用顺序比较的方法。首先比较L[i]和L[i-1]，如果L[i-1]≤ L[i]，则L[1..i]已排好序，第i遍处理就结束了；否则交换L[i]与L[i-1]的位置，继续比较L[i-1]和L[i-2]，直到找到某一个位置j(1≤j≤i-1)，使得L[j] ≤L[j+1]时为止。图1演示了对4个元素进行插入排序的过程，共需要(a),(b),(c)三次插入。 
> 直接插入排序是稳定的。算法时间复杂度是O(n ^2) 。


```
func insertSort(arr []int){
   for i:=0;i<len(arr);i++{
      tmp:=arr[i]
      key := i-1
      for key>=0 && tmp<arr[key]{
         arr[key+1] = arr[key]
      }
      if key+1!=i{
        arr[key+1] = tmp
      }
   }
}
```

# golang sort package

type sortSlice []int

func(s sortSlice) Len() int{
   return s
}

func(s sortSlice)Less(i, j int)bool {
   return s[i]>s[j]
}

func (s sortSclice) Swap(i，j) {
  s[i],s[j]=s[j],s[i]
}










