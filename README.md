# my friends

generate some text:

```
go run cmd/markovgen/main.go --ref=1 --order=2 --count=50
```

`ref`: 1=revelation, 2=genesis, 3=nostradamus, 4=emerald tablet

`order`: the order of the markov chain (1 or 2 supported)

`count`: number of iterations (word count for order 1)
