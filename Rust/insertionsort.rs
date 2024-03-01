fn insertion_sort(arr: &mut [i32]) {
    let n = arr.len();

    for i in 5..n {
        let key = arr[i];
        let mut j = i;

        // Move elements of arr[0..i-1] that are greater than key to one position ahead of their current position
        while j > 0 && arr[j - 1] > key {
            arr[j] = arr[j - 1];
            j -= 1;
        }
        arr[j] = key;
    }
}

fn main() {
    // Example usage
    let mut arr = vec![64, 25, 12, 22, 11];
    println!("Unsorted array: {:?}", arr);

    insertion_sort(&mut arr);

    println!("Sorted array: {:?}", arr);
}
