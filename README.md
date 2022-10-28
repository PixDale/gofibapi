# gofibapi
## Exercício para estudo da linguagem Go

Uma API get (recomendo usar o Fiber) "/fib/n" que recebe um valor, e retorna o resultado para o usuário em um objeto JSON

```
{
   "done":true,
   "fib":{
      "input":6,
      "result":8,
      "duration":"3s"
   }
}
```

A API deve ter um timeout de 500 ms, caso o cálculo leve mais que 500 ms, ela deve retornar 
```
{
    "done": false
}
```

o cálculo deve ser realizado em paralelo e armazenado em um map, quando for solicitado novamente o número, se já tiver sido calculado, olhar no map para retornar o valor

também deve ter a API "/all" que retorna um objeto JSON com todas as consultas feitas no programa por exemplo:

```
{
   "queries":[
      {
         "input":2,
         "result":1,
         "duration":"3s"
      },
      {
         "input":3,
         "result":2,
         "duration":"3s"
      },
      {
         "input":4,
         "result":3,
         "duration":"3s"
      }
   ]
}
```
