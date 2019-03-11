module crypto-pals

require (
	example.com/set1 v0.0.0
	example.com/utils v0.0.0
)

replace example.com/utils => ./utils

replace example.com/set1 => ./set1
