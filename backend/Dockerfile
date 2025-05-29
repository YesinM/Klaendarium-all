#komentarze są tylko moim konspektem, dlatego można usunąć :)
#from język:vers-"system" obrazu
FROM golang

#folder *wirtualny*
WORKDIR /app
#kopiowanie plików do utworzonego folderu
COPY . . 
#RUN polecenie
RUN go mod download
#/backend -copy to> /app
#

#RUN działa tylko podczas ładowania kontenera, dlatego RUN go run main.go stworzy "martwy" kontener
CMD ["go", "run", "main.go"]