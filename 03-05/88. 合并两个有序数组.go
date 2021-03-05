package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	// index可以直接用lr表示
	// l,r:=m-1,n-1
	// for l>=0 || r>=0{
	// 	if l>=0 && r>=0{
	// 		if nums1[l]>=nums2[r]{
	// 			nums1[l+r+1] = nums1[l]
	// 			l--
	// 		}else{
	// 			nums1[l+r+1] = nums2[r]
	// 			r--
	// 		}
	// 	}else if r>=0{
	// 		nums1[l+r+1] = nums2[r]
	// 		r--
	// 	}else{
	// 		break
	// 	}
	// }
	// 为了双百 l r也不要了
	// ????? 这个空间不是100woc
	for m-1 >= 0 || n-1 >= 0 {
		if m-1 >= 0 && n-1 >= 0 {
			if nums1[m-1] >= nums2[n-1] {
				nums1[m+n-1] = nums1[m-1]
				m--
			} else {
				nums1[m+n-1] = nums2[n-1]
				n--
			}
		} else if n-1 >= 0 {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			break
		}
	}
	// 先将nums1保存在其他地方
	// tmp:=make([]int,m)
	// copy(tmp,nums1)
	// l,r:=0,0
	// for l<m || r<n{
	// 	if l<m && r<n{
	// 		if tmp[l] <= nums2[r]{
	// 			nums1[l+r] = tmp[l]
	// 			l++
	// 		}else {
	// 			nums1[l+r] = nums2[r]
	// 			r++
	// 		}
	// 	}else if l<m{
	// 		nums1[l+r] = tmp[l]
	// 		l++
	// 	}else if r<n{
	// 		nums1[l+r] = nums2[r]
	// 		r++
	// 	}
	// }
}
