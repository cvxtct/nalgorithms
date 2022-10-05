"""Regex practice"""
import re

# match = re.search(pattern, str)
# re.sub(pat, replacement, str) -- returns new string with all replacements,
# a, X, 9, < -- ordinary characters just match themselves exactly. The meta-characters which do not match themselves because they have special meanings are: . ^ $ * + ? { [ ] \ | ( ) (details below)
# . (a period) -- matches any single character except newline '\n'
# \w -- (lowercase w) matches a "word" character: a letter or digit or underbar [a-zA-Z0-9_]. Note that although "word" is the mnemonic for this, it only matches a single word char, not a whole word. \W (upper case W) matches any non-word character.
# \b -- boundary between word and non-word
# \s -- (lowercase s) matches a single whitespace character -- space, newline, return, tab, form [ \n\r\t\f]. \S (upper case S) matches any non-whitespace character.
# \t, \n, \r -- tab, newline, return
# \d -- decimal digit [0-9] (some older regex utilities do not support \d, but they all support \w and \s)
# ^ = start, $ = end -- match the start or end of the string
# \ -- inhibit the "specialness" of a character. So, for example, use \. to match a period or \\ to match a slash. If you are unsure if a character has special meaning, such as '@', you can try putting a slash in front of it, \@. If its not a valid escape sequence, like \c, your python program will halt with an error.
# 001 A very simple solution:

str = 'an example word:cat!!'
match = re.search(r'\sword:\w\w\w', str)
if match:
    print(match.group())
match = re.search(r'iii', 'piiig')
if match:
    print(match.group())
match = re.search(r'\d\w\d', 'pi1i2ig')
if match:
    print(match.group())
match = re.search(r'\d+\w\d', 'pi12i2ig')
if match:
    print(match.group())
match = re.search(r'^\d+\w\d', 'pi12i2ig')
if match:
    print(match.group())
str = 'purple alice-b@google.com monkey dishwasher'
match = re.search(r'\w+@\w+', str)
if match:
    print(match.group()) 
str = 'purple al.ice-b@goo-gle.com monkey dishwasher'
match = re.search(r'[\w.-]+@[\w.-]+', str)
if match:
    print(match.group())  
str = 'purple al.123ice-b@goo-gle.com monkey dishwasher'
match = re.search(r'[\w.-]+@[\w.-]+', str)
if match:
    print(match.group())
str = 'purple al.123ice-b@goo-gle.com monkey dishwasher'
match = re.search(r'[0-9]+', str)
if match:
    print(match.group())  
str = 'purple al.123ice-b@goo-gle.com monkey dishwasher'
match = re.findall('g', str)
if match:
    print(match)  
str = 'purple al.123ice-b@goo-gle.com monkey dishwasher'
match = re.search(r'[a-z,\s.0-9]+', str)
if match:
    print(match.group()) 



strings = ['abcNdgM', 'abcdg', 'MrtyNNgMM']

for s in strings:
    # Repeat the cycles of transforming M/N with previous or subsequent characters:
    while True:
        s_new = re.sub(r'N.', '', re.sub(r'(.)M', r'\1\1', s))
        # print('s_new:', s_new)
        if s == s_new:
            break
        s = s_new
    # Remove any remaining Ms and Ns:
    s = re.sub(r'[MN]+', '', s)
    print(s)


inp = 'Today is the greatest day ever'
l = inp.split()
helper = []
for i in l:
    helper.append(len(i))

greatest = max(helper)
#print(greatest)

for i in l:
    if len(i) == greatest:
        print(i)

l = inp.split()
l_dict = dict()
for i in l:
    for j in l:
        if i.count(j) > 1:
            l_dict[i] = i.count(j)
        if len(l_dict) != 0:
            for i in l_dict:
                print(i)
            
            break
