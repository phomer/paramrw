# Parameterized Random Walks

A simple example to play around with parameterizing random walks. Are there basic control parameters we can add to change the behaviour? Color and animation were addded later.

Initial starting point:

> The idea is to try to make classic random walks faster. I think what I am suggesting could be neat and it could be silly.
> 
> A classic random walk goes like this: At each stage it flips a coin and moves to some adjacent state randomly. The trouble is that for a linear walk this takes time order n2 to move about n steps. The idea is to try and speed this up. The simplest idea is to remember all states that the random walk has already visited. Do this via some simple caching method.
> 
> Then a step is this: Flip a coin and consider the move from the
> current state S to a new one S′. If S′ is not already in the cache, make the move as before. Of course update the cache to include the new state S′.
> 
> If S′ is already visited that is it is in the cache, then flip another coin and try another move. If these moves keep hitting a old state eventually make the move anyway.
> 
> Simple point: Now the line-dimensional walk moves n steps in n steps. Do you see why? What happen in the two dimensional case is thus the first interesting question. I cannot see the behavior yet so I hope that you might try and implement the method and see how it behaves.
> 
> The experiment could first try simple walk in two space on the lattice (i, j). The “eventually” could be a fixed limit that is some parameter. Then a run will keep track of how far the walk gets in n steps and perhaps how often it does the “eventually” step. A possible cool idea might be to show the lattice picture of the states hit after n steps.
> 
> Dick
> 

