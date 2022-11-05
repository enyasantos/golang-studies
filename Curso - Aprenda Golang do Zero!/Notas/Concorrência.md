<h1>Concorrência</h1>

    - Uma das maiores forças da linguagem
    - Concorrência != Paralelismo

<p align="center">
    <img src="https://miro.medium.com/max/744/1*kIpwwNMEzEMiknbAVlUZyg.jpeg" 
    />
</p>



<h2>Goroutines</h2>
<p>
    Goroutines são funções ou métodos executados em concorrência. Podemos pensar nelas como uma especie de lightweight thread que são gerenciadas pelo runtime do Go.
</p>

<p>
    Chamamos de lightweight thread pois o custo para sua criação é muito menor quando comparada com um thread de verdade. Outro ponto positivo é que o runtime consegue aumentar ou diminuir a quantidade de goroutines de acordo com a necessidade da aplicação, enquanto o número de thread normalmente é fixo.
</p>

```
    go <medoto>
```

<h2>WaitGroup</h2>
<p>
    - Sincronizar Goroutines
    - Usa um "contador"
</p>

```
    waitGroup.Add(4)

    waitGroup.Done()

    waitGroup.Wait()
```

<h2>Canais</h2>
<p>
    - Sincronizar Goroutines
    - Canal de comunicação para enviar e receber dados
    - Enviar e Receber, essas operações bloqueiam a execução do programa
    - Verificar se o canal está aberto ou fechado
</p>

