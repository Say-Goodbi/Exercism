module Luhn (isValid) where

import Data.Char

operate :: [Int] -> Int
operate (x:z:xs)
  | y > 9 = x + y - 9 + operate xs
  | otherwise = x + y + operate xs
  where y = z * 2
operate t = sum t

isValid :: String -> Bool
isValid n
  | length trimmed > 1 && all isDigit trimmed =(==0).(`mod` 10).operate.map digitToInt.reverse $ trimmed
  | otherwise = False
  where trimmed = filter (/=' ') n