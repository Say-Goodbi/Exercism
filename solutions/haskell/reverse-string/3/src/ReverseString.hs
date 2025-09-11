module ReverseString (reverseString) where

reverseString :: String -> String
reverseString [] = ""
reverseString s = last s : reverseString (take (length s - 1) s)