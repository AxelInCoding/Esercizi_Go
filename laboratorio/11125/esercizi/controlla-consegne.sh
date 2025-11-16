#!/bin/bash

#https://stackoverflow.com/questions/58637568/for-loop-in-linux-treats-pattern-as-filename-when-no-files-exist
shopt -s nullglob # expands the glob to empty string when there are no matching files in the directory.

# PREDISPOSIZIONE PER USARE QUESTO SCRIPT:
# Creare una propria dir di lavoro (la chiamerò prog1, ma potete chiamarla come volete) 
# In prog1 salvare e rendere eseguibile (chmod u+x) il file controlla-consegne.sh 
# In prog1 creare una dir Lab<numLab> (due cifre)
# In prog1 salvare il file degli esercizi (Lab<numLab>.md)
# Creare in Lab<numLab> una sottodir consegne
# Salvare in Lab<numLab> la dir bbtest di quel lab
# Lanciare il comando controlla-consegne.sh seguito da  <numLab> (due cifre)

# struttura di file e dir occorrenti per eseguire lo script (esempio su Lab02)
# |--controlla-consegne.sh
# |--Lab02
# |  |--bbtest
# |  |--consegne
# |--Lab02.md

####################################################################################
# verifica che venga passato un parametro
if
    test "$#" -ne 1
then
    echo "Usage: $0 <numero (di due cifre) del laboratorio da controllare>"
    exit 1
fi

SUPPRESS="--suppress-common-lines"
FILESPEC="$PWD/Lab${1}.md"
DIRLAB="$PWD/Lab${1}"
DIRCONSEGNE="$DIRLAB/consegne"
DIRBBTEST="$DIRLAB/bbtest"
LOG="Lab${1}_controlla-consegne.log"
INITIALDIR="$PWD/"

####################################################################################
# verifica che file e dir esistano
if
    ! test -s "$FILESPEC" || \
    ! test -d "$DIRLAB" || \
    ! test -d "$DIRCONSEGNE" || \
    ! test -d "$DIRBBTEST"
then
    echo "Manca qualche file/dir..."
    exit 2
fi

####################################################################################
ELENCOESERCIZI=$(grep nomefile "$FILESPEC"|cut -f2 -d:)
#TOTESERCIZI=$(echo $ELENCOESERCIZI|tr " " "\n"|wc -l)
ESARRAY=( $ELENCOESERCIZI )
TOTESERCIZI=${#ESARRAY[@]}
echo -e "Esercizi da controllare:\n$ELENCOESERCIZI ($TOTESERCIZI)" | tee "$LOG"     # LOG globale

####################################################################################
# scendo in dir consegne e cerco esercizi
find "$DIRCONSEGNE" -type d | while
 read dir
do
 #realdir=$(realpath --relative-to . "$dir")
 realdir=${dir/$INITIALDIR/}
 echo "/// directory: $realdir"

 pushd "$dir" >/dev/null || continue

 for es in $ELENCOESERCIZI
 do
    # c'è?
    if
        ! test -s "$es"
    then
        echo "!!! MANCA: $es"
        continue
    fi

    (( FALLITI = 0 ))
    (( NRTEST = 0 ))

    echo "... compilo '$es' (e stampo eventuali errori di compilazione)"
    go build "$es" | tee "$es.compile"
    if
        ! test -s "$es.compile"
    then
        rm "$es.compile"

        echo  "... eseguo i/il test e stampo eventuali errori" # (no output = OK o non erano presenti test)"
        EXENAME=$(basename "$es" .go)
        for testinput in "$DIRBBTEST"/"$EXENAME".*.input; do
            #echo $testinput

            expected=${testinput//input/expected}
            #echo $expected
            #ls -l

            # fa il `diff` tra l'output dell'esecuzione del programma (con redirezione input) e l'out expected
            #realexpected=$(realpath --relative-to . "$expected")
            realexpected=${expected/$INITIALDIR/}
            #echo PWD: $PWD

            (( NRTEST = NRTEST + 1))

            #echo "... diff (studente - expected)"
            diff "$SUPPRESS" -y -W 180 <(timeout 1 "./$EXENAME" < "$testinput"|head -n 100) "$expected" | sed 's/^/\t/'
            # per testare successo della pipe
            if [ "${PIPESTATUS[0]}" -eq 0 ]; then
                echo -e "\n... ${PWD/$DIRCONSEGNE\//}/$EXENAME CORRECT on $realexpected"
            else
                echo -e "\n... ${PWD/$DIRCONSEGNE\//}/$EXENAME DIFFERS on $realexpected"
                (( FALLITI = FALLITI + 1))
            fi
        done #| tee $es.run

        echo "... FALLITI: $es,$FALLITI,$NRTEST"
        
        # rimuove eseguibile (non serve più)
        rm "$EXENAME"
    fi
 done | tee "$LOG"    # locale
 echo '########################################
 Significato dei simboli in diff
 (spazio) → le due righe sono identiche
 | → le due righe differiscono (entrambe presenti)
 < → riga presente solo nel primo output (sinistra)
 > → riga presente solo nel secondo output (destra)
 \ → riga del primo output troncata o incompleta
 / → riga del secondo output troncata o incompleta
########################################' | tee -a $LOG

 NONCONSEGNATI=$(grep -c MANCA "$LOG")
 (( CONSEGNATI=TOTESERCIZI-NONCONSEGNATI ))
 echo "### consegnati: $CONSEGNATI su $TOTESERCIZI" | tee -a "$LOG"

 echo "### discrepanze nomi (studente - expected)"| tee -a "$LOG"
 diff -y --suppress-common-lines <(ls *.go|sort)  <(echo $ELENCOESERCIZI|tr " " "\n"|sort) | tee -a "$LOG"

 popd >/dev/null || continue
done | tee -a "$LOG"     # LOG globale