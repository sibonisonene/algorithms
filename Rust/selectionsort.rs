fn selection_sort(arr: &mut [i32]) {
    let n = arr.len();

    for i in 0..n - 1 {
        let mut min_index = i;

        // Find the index of the minimum element in the remaining unsorted part
        for j in i + 1..n {
            if arr[j] < arr[min_index] {
                min_index = j;
            }
        }

        // Swap the found minimum element with the first element
        arr.swap(i, min_index);
    }
}

fn main() {
    // Example usage
    let mut arr = vec![64, 25, 12, 22, 11];
    println!("Unsorted array: {:?}", arr);

    selection_sort(&mut arr);

    println!("Sorted array: {:?}", arr);
}
