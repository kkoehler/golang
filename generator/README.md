# AVRO Generierung mit Go

Zuerst müssen Sie das Tool GoGen-Avro installieren:

```
go get github.com/actgardner/gogen-avro/cmd/gogen-avro
```

Danach können Sie aus dem AVRO Schema Code generieren

```
go generate cmd/*.go
```

... und jetzt die Anwendung ausführen:

```
go run cmd/main.go
```

