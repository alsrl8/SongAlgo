import sys
import heapq

N = int(sys.stdin.readline())

max_heap = []
min_heap = []
for _ in range(N):
    num = int(sys.stdin.readline())
    arr = [num]
    if len(max_heap) > len(min_heap):
        if -max_heap[0] > num:
            temp = heapq.heappop(max_heap)
            heapq.heappush(max_heap, -num)
            heapq.heappush(min_heap, -temp)
        else:
            heapq.heappush(min_heap, num)
    elif len(max_heap) == len(min_heap):
        if len(max_heap) == 0:
            heapq.heappush(max_heap, -num)
        else:
            arr = [
                num,
                -heapq.heappop(max_heap),
                heapq.heappop(min_heap),
            ]
            arr.sort()
            heapq.heappush(max_heap, -arr[0])
            heapq.heappush(max_heap, -arr[1])
            heapq.heappush(min_heap, arr[2])
    print(-max_heap[0])
