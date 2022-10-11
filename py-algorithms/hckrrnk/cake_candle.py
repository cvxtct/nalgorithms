def birthdayCakeCandles(candles):
    # Write your code here
    max_height = max(candles)
    print(max_height)
    num = candles.count(max_height)
    return num

if __name__ == '__main__':
    print(birthdayCakeCandles([3,2,1,3]))
