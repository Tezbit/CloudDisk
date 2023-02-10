set path=/D/Works/JetBrains/GoProject/MsgSystem/proto
set name=product.proto
set code=0D23CF3A7AB3ED75
 docker run --rm -v %path%:%path% -w %path% -e ICODE=%code% cap1573/cap-protoc -I %path% --go_out=%path% %path%/%name% --micro_out=%path%
