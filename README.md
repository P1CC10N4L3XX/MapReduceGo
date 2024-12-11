# MapReduceGo

Si vuole implementare un sistema distribuito che esegua le operazioni di map e reduce in linguaggio di programmazione GO. 

# Specifiche tecniche

- Sulla macchina deve essere istallato l'IDL compiler protoc
- Sono supportati sistemi operativi Windows, Linux, MACOS

# Utilizzo del programma

- Nella directory principale del progetto eseguire il comando 

```
make

```
- Eseguire il comando per aggiungere all'utente i permessi di esecuzione sul file ```init.sh```

```
chmod -x init.sh

```
- Eseguire il comando per avviare i vari mapper e reducer

```
./init.sh

```
- Eseguire il comando per avviare il master

```
./builds/master <input_file>

```

