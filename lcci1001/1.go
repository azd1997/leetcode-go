package lcci1001

// 思路：
// 1.最笨的方法： 将B追击到A末尾，对A重新排序 O(nlogn)
// 2.准备一个备份数组，用双指针指向A，B头部，从左向右开始取指针处元素比较，小者入备份数组此时位置 O(n)/O(n)
// 3. 由于A空间足够，因此，从右向左以双指针方式填充数据可不需要额外空间。 O(n)/O(1)

// 双指针 + A原地合并
func merge(A []int, m int, B []int, n int)  {
	// 边界条件
	if n == 0 {return}
	if m == 0 {
		for i:=0; i<n; i++ {A[i] = B[i]}
		return
	}

	// 双指针
	p := m+n-1
	pa := m-1
	pb := n-1
	for pa >= 0 && pb >= 0 {    // 有一边遍历完了就停止
		if A[pa] >= B[pb] {
			A[p] = A[pa]
			pa--; p--
		} else {
			A[p] = B[pb]
			pb--; p--
		}
	}

	//  现在要检查是哪个先遍历完，如果是B先遍历完则剩余的A不需要动；
	// 反之则需要将B剩余赋值到A剩余位置
	if pb >= 0 {
		for pb >= 0  {
			A[pb] = B[pb]
			pb--
		}
	}
}
