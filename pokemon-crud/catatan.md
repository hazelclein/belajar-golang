note :

### respon

### Success Response:
```
{
  "status": "success",
  "message": "Pokemon berhasil disimpan",
  "data": {
    "id": 1,
    "name": "bulbasaur",
    "height": 7,
    "weight": 69,
    "base_experience": 64,
    "types": "grass, poison"
  }
}
```

### Error
```
{
  "status": "error",
  "message": "Validasi gagal",
  "error": "Pokemon ID tidak boleh kosong"
}
```

### penting
```
- ID ga boleh kosong atau 0
- Nama minimal 2 karakter
- Height, Weight, Base Experience ga boleh negatif
- Types ga boleh kosong
```

### Test validasi

### ID kosong
```
POST http://localhost:8080/pokemon/fetch/0

Hasil :
{
     "status": "error",
     "message": "Validasi gagal",
     "error": "Pokemon ID harus lebih dari 0"
   }

```
### update data invalid
```
PUT http://localhost:8080/pokemon/1

Body teks :
{
     "name": "A",
     "height": -5,
     "weight": 100,
     "base_experience": 64,
     "types": "grass"
   }

Hasil :
{
     "status": "error",
     "message": "Validasi data gagal",
     "error": "Nama pokemon minimal 2 karakter"
   }
```

### endpoint
```
- POST /pokemon/fetch/{id} - ambil datapPokemon dari PokeAPI + sv ke database

- GET /pokemon - ambil all pokemon dari database

- GET /pokemon/{id} - Ambil Pokemon berdasarkan ID

- GET /pokemon/search?name={name} - cari pokemon pake nama

- GET /pokemon/type/{tipe} - cari pokemon pake tipe

- okkey lanjutGET /pokemon?page=1&limit={limit} - ambil data bedasarkan limit

- PUT /pokemon/{id} - update Pokemon

- DELETE /pokemon/{id} - hapus Pokemon
```


