module Prime (nth) where

primes :: [Integer] -> [Integer]
primes [] = []
primes (p:xs) = p : primes [x | x <- xs, x `mod` p /= 0]

nth :: Int -> Maybe Integer
nth n
  | n > 0 = Just $ (0:primes [2..]) !! n
  | otherwise = Nothing