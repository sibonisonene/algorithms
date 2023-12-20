def insertion_sort(arr)
    n = arr.length
  
    (1...n).each do |i|
      key = arr[i]
      j = i - 1
  
      # Move elements of arr[0..i-1] that are greater than key to one position ahead of their current position
      while j >= 0 && arr[j] > key
        arr[j + 1] = arr[j]
        j -= 1
      end
  
      arr[j + 1] = key
    end
  end
  
  # Example usage
  arr = [64, 25, 12, 22, 11]
  puts "Unsorted array: #{arr}"
  
  insertion_sort(arr)
  
  puts "Sorted array: #{arr}"
  