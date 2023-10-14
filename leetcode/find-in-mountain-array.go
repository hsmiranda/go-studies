/**
    https://leetcode.com/problems/find-in-mountain-array/description/
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, arr *MountainArray) int {
    peak := findPeak(arr);

    ans := findLeft(arr, peak, target);
    if ans != -1 {
        return ans;
    }

    return findRight(arr, peak, target);
}

func findRight (arr *MountainArray, peak int, target int) int {
    left := peak+1;
    right := arr.length()-1;

    for left <= right {
        middle := left + (right-left)/2;
        if arr.get(middle) == target {
            return middle;
        }

        if arr.get(middle) < target {
            right = middle-1;
        } else {
            left = middle+1;
        }
    }
    return -1;
}

func findLeft (arr *MountainArray, peak int, target int) int {
    left := 0;
    right := peak;

    for left <= right {
        middle := left + (right-left)/2;

        if arr.get(middle) == target {
            return middle;
        } 

        if arr.get(middle) < target {
            left = middle+1;
        } else {
            right = middle-1;
        }
    }

    return -1;
}

func findPeak(arr *MountainArray) int {
    left := 1;
    right := arr.length()-2;
    for left <= right {
        middle := left+(right-left)/2;

        if arr.get(middle) < arr.get(middle+1) {
            left = middle+1;
            continue;
        }

        if arr.get(middle) < arr.get(middle-1) {
            right = middle-1;
            continue;
        }

        return middle;
    }

    return -1;
}
