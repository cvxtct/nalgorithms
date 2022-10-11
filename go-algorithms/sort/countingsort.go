package sort

// Counting sort is a sorting technique based on keys between a specific range.
// It works by counting the number of objects having distinct key values (kind of hashing).
// Then do some arithmetic to calculate the position of each object in the output sequence.
// https://www.youtube.com/watch?v=OKd534EWcdk
// O (n+k) n: number of items we have in the original array, k: the numbers we could have for each item in the array.
