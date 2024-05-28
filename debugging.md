# Debugging commands

## Create 2 blocks simultaneously
```shell
(
  curl -X POST http://127.0.0.1:8080/blocks -H "Content-Type: application/json" -d '{"data": "VLvvaUvQHt",
  "hash": "4db150153938abef026b0c327d43097f1e5480b6205e6fc1ad8741df061927aa"}' &
  curl -X POST http://127.0.0.1:8080/blocks -H "Content-Type: application/json" -d '{"data": "XAGGuWCvGq",
  "hash": "dc751acf4f4694e08c66e497b9e320d11f0fcdac20d5709f400cad4ace092ce0"}' &
  wait
)
```
```shell
(
  curl -X POST http://127.0.0.1:8000/generate_block  &
  curl -X POST http://127.0.0.1:8000/generate_block  &
  wait
)
```