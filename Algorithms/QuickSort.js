function Partition(arr) {
    // get pivot and remove from the array
    const pivot = arr.pop();
    // allocate our smaller and bigger arrays
    const smaller = []
    const bigger = []

    // iterate over our array (without the pivot) to sort into {smaller} and {bigger}
    // items smaller or equal to the pivot go into {smaller} and items bigger go into {bigger}
    arr.forEach(elem => {
        if (elem <= pivot) {
            smaller.push(elem)
        } else {
            bigger.push(elem)
        }
    })

    // return our 3 items
    return {
        pivot,
        smaller,
        bigger
    }
}

function QuickSort(arr) {
    // if our array is less than 2 items, it's sorted!
    if (arr.length < 2) {
        return arr
    }

    // else, everything with more than 2 items can be broken down
    const {
        pivot,
        smaller,
        bigger
    } = Partition(arr)

    // and our *sorted* smaller bucket, + our pivot, and our *sorted* bigger bucket will give us the whole array
    return [ ...QuickSort(smaller), pivot, ...QuickSort(bigger) ]
}

const arr = [5, 2, 1, 6, 13, -3];

console.log(QuickSort(arr))