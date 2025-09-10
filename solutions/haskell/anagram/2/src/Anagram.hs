module Anagram (anagramsFor) where
import Data.Char

delete :: Char -> [Char] -> [Char]
delete _ [] = []
delete t (x:xs)
  | t == x = xs
  | otherwise = x:delete t xs

sort :: [Char] -> [Char]
sort [] = []
sort xs = m : (sort . delete m) xs
  where m = minimum xs

isAnagram :: String -> String -> Bool
isAnagram t x
  | t == x = False
  | sort t == sort x = True
  | otherwise = False

anagramsFor :: String -> [String] -> [String]
anagramsFor xs xss = filter (isAnagram t . map toLower) xss
  where t = map toLower xs