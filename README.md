# MapReduceGo

Si vuole implementare un sistema distribuito che esegua le operazioni di map e reduce in linguaggio di programmazione GO. 

# Specifiche tecniche

- Sulla macchina deve essere istallato l'IDL compiler protoc
- Sono supportati sistemi operativi Windows, Linux, MACOS

# Utilizzo del programma

- Nella directory principale del progetto eseguire il comando di seguito 

```
make

```
Eseguire il comando di seguito per aggiungere all'utente i permessi di esecuzione sul file ```init.sh```

```
chmod +x init.sh

```
Eseguire il comando di seguito per avviare i vari mapper e reducer

```
./init.sh

```
Eseguire il comando di seguito per avviare il master (assicurandosi che i server mapper e reduce siano stati avviati)

```
./builds/master <input_file>

```

Nella directory del progetto si può trovare un file ```prova.txt``` che definisce come devono essere strutturati i dataset da passare al master

# File di output

Sarà possibile trovare i file di output ```reducer1.txt``` e ```reducer2.txt``` all'interno della cartella ```/output```

