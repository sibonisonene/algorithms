def selection_sort(arr)
    n = arr.length
  
    (0...n - 1).each do |i|
      min_index = i
  
      # Find the index of the minimum element in the remaining unsorted part
      (i + 1...n).each do |j|
        min_index = j if arr[j] < arr[min_index]
      end
  
      # Swap the found minimum element with the first element
      arr[i], arr[min_index] = arr[min_index], arr[i]
    end
  end
  
  # Example usage
  arr = [64, 25, 12, 22, 11]
  puts "Unsorted array: #{arr}"
  
  selection_sort(arr)
  
  puts "Sorted array: #{arr}"
  