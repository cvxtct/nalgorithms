import re
import operator

from collections import Counter
def most_common_char(strng):
  s = sorted(strng)


  frequency = Counter(list(s))
  for k, v in frequency.most_common(3):
    print(k, v)

if __name__ == '__main__':
    string = "aabbbccde"
    print(most_common_char(string))