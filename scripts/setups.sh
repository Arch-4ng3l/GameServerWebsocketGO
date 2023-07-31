#!/bin/bash


username=$PSQLUN
password=$PSQLPW


function clean() {

    sed -e "s/\${1}/${username}/"  clear.sql > tmp.sql
    tmp="./tmp.sql"
    sudo -i -u postgres psql -q < $tmp
    sudo rm $tmp

}

function setup() {
    sed -e "s/\${1}/${username}/" -e "s/\${2}/${password}/" setup.sql > tmp.sql
    tmp="./tmp.sql"
    sudo -i -u postgres psql -q < $tmp
    sudo rm $tmp
}

function showUsage() {
    echo "-c clear      Clear the Database"
    echo "-s setup      Create User, Table and Database"
}

cleanF=false
setupF=false
while getopts "csh" flag
do
    case "${flag}" in
        c) cleanF=true;;
        s) setupF=true;;
        h) showUsage;;
        *)
    esac

done


if [ $cleanF = true ]; then
    echo "[*] Cleaning Database"
    clean
    echo "[+] Done Cleaning"
fi

if [ $setupF = true ]; then
    echo "[*] Setting Up Database"
    setup
    echo "[+] Done Setting Up"
fi

