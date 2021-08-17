function SelectionSort(arr) {
    // if the array is less than 2 elements, it's already sorted
    if (arr.length < 2) return arr;

    // loop over the array
    for (let i = 0; i < arr.length; i++) {
        // index of the minimum element going through this loop
        let minIdx = i;
        // lookahead to check and see for potential smaller indexes
        let lookahead = minIdx + 1;

        // while lookahead hasn't gotten to the end, increment it
        while (lookahead < arr.length) {
            // if lookahead's element is smaller
            if (arr[lookahead] < arr[minIdx]) {
                // update our min index
                minIdx = lookahead
            }

            lookahead++
        }

        // swap our current element with the minimum element we found
        const temp = arr[i];
        arr[i] = arr[minIdx];
        arr[minIdx] = temp;
    }

    return arr
}

const arr = [5, 2, 1, 6, 13, -3];

console.log(SelectionSort(arr))

