<a name="readme-top"></a>  
[![English](https://img.shields.io/badge/lang-en-blue.svg)](README-en.md) [![Italiano](https://img.shields.io/badge/lang-it-blue.svg)](README.md) ![License](https://img.shields.io/github/license/anond0rf/vecchioposter) [![GitHub Release](https://img.shields.io/github/v/release/anond0rf/vecchioposter?label=release)](https://github.com/anond0rf/vecchioposter/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/anond0rf/vecchioposter)](https://goreportcard.com/report/github.com/anond0rf/vecchioposter) [![Go Version](https://img.shields.io/github/go-mod/go-version/anond0rf/vecchioposter)](https://github.com/anond0rf/vecchioposter)  

<br />
<div align="center">
  <a href="https://github.com/anond0rf/vecchioposter">
    <img src="logo.png" alt="Logo" width="80" height="80">
  </a>
<h3 align="center">VecchioPoster</h3>
  <p align="center">
    <strong>VecchioPoster</strong> è un'applicazione a riga di comando per postare su <a href="https://vecchiochan.com/">vecchiochan.com</a>.
    <br />
    <br />
    <a href="#download"><strong>Inizia »</strong></a>
    <br />
    <br />
    <a href="https://github.com/anond0rf/vecchioposter/releases">Release</a>
    ·
    <a href="https://github.com/anond0rf/vecchioposter/issues">Segnala Bug</a>
    ·
    <a href="https://github.com/anond0rf/vecchioposter/issues">Richiedi Feature</a>
  </p>
</div>

## Caratteristiche

VecchioPoster consente di postare su [vecchiochan](https://vecchiochan.com) via riga di comando.  
L'applicazione astrae i dettagli dell'invio del form e della gestione delle richieste utilizzando [vecchioclient](https://github.com/anond0rf/vecchioclient).
Con i comandi e le opzioni disponibili, puoi:

- Creare nuovi thread su board specifiche
- Rispondere a thread esistenti

Supporta l'header `User-Agent` personalizzato per il client interno e il logging dettagliato (vedi [Utilizzo](#utilizzo)).

## Indice

1. [Download](#download)
2. [Utilizzo](#utilizzo)
   - [Pubblicare un nuovo thread](#pubblicare-un-nuovo-thread)
   - [Pubblicare una risposta](#pubblicare-una-risposta)
3. [Compilare il progetto](#compilare-il-progetto)
4. [Licenza](#licenza)

## Download

VecchioPoster è disponibile per Windows, GNU/Linux e MacOS.  
L'eseguibile dell'ultima versione si può scaricare da [qui](https://github.com/anond0rf/vecchioposter/releases).

## Utilizzo

Di seguito alcuni esempi su come utilizzare l'applicazione.  
Per semplicità, si assume che `vecchioposter` sia il nome dell'eseguibile.
Consulta l'opzione `--help` (`-h`) per maggiori dettagli.

```sh
vecchioposter -h
```

- #### Pubblicare un nuovo thread

  Per creare un nuovo thread, usa il comando `new-thread`:

  ```sh
  vecchioposter new-thread -b b -B "Questo è un nuovo thread sulla board /b/" -f path/to/file.jpg
  ```

  `--board` (`-b`) è l'unico flag **obbligatorio**, ma tieni presente che, poiché le impostazioni variano tra le board, potrebbero essere richiesti più flag per postare (ad esempio, non è possibile postare un nuovo thread senza embed né file su /b/).  
  Per l'elenco completo dei flag, con abbreviazioni e alias dei comandi, esegui:

  ```sh
  vecchioposter new-thread -h
  ```

- #### Pubblicare una risposta

  Per rispondere a un thread esistente, usa il comando `post-reply`:

  ```sh
  vecchioposter post-reply -b b -t 1 -B "Questa è una risposta al thread #1 sulla board /b/" -f path/to/file1.mp4 -f path/to/file2.webm
  ```

  `--board` (`-b`) e `--thread` (`-t`) sono i soli flag **obbligatori**, ma considera che, poiché le impostazioni variano tra le board, potrebbero essere necessari più flag per postare.  
  Per l'elenco completo dei flag, con abbreviazioni e alias dei comandi, esegui:

  ```sh
  vecchioposter post-reply -h
  ```

Informazioni aggiuntive:
 - Per semplicità d'uso, viene fornito il flag `--msg-file` (`-m`) per impostare il corpo del messaggio a partire da un file.

    ```sh
    vecchioposter post-reply -b b -t 1 -m path/to/msg.txt
    ```  
    Con questo comando viene letto il contenuto di `msg.txt` e impostato come corpo del messaggio.  
    Questo flag sostituisce `--body` (`-B`), quindi è possibile utilizzare uno solo dei due.

- Per impostare un header `User-Agent` personalizzato da usare nel client interno, è disponibile il flag `--user-agent` (`-u`):

    ```sh
    vecchioposter new-thread -u "CustomUserAgent" -b b -B "Questo è un nuovo thread sulla board /b/" -f path/to/file.jpg
    ```  

- È possibile abilitare il logging dettagliato con il flag `--verbose` (`-v`):

    ```sh
    vecchioposter new-thread -v -b b -B "Questo è un nuovo thread sulla board /b/" -f path/to/file.jpg
    ```  

## Compilare il progetto

Per compilare VecchioPoster dal codice sorgente:

1. Assicurati di avere installato [Go](https://golang.org/dl/).
2. Clona il repository con [git](https://github.com/git/git):

   ```sh
   git clone https://github.com/anond0rf/vecchioposter.git
   ```

2. Spostati nella directory del progetto:

   ```sh
   cd vecchioposter
   ```

3. Compila il progetto:

   ```sh
   go build
   ```

Verrà generato un file eseguibile nella directory principale del progetto.

## Licenza

VecchioPoster è concesso in licenza sotto la [Licenza LGPL-3.0](https://www.gnu.org/licenses/lgpl-3.0.html).

Questo significa che puoi usare, modificare e distribuire il software, a condizione che eventuali versioni modificate siano anch'esse concesse in licenza sotto la LGPL-3.0.

Per maggiori dettagli, consulta il testo completo della licenza nel file [LICENSE](./LICENSE).

Copyright © 2024 anond0rf