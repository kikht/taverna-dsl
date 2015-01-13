DSL_PATH=$GOPATH/src/gitlab.ict.sbras.ru/taverna/taverna-dsl
pushd "$DSL_PATH/t2compile"
go build
popd

pushd "$DSL_PATH/t2extract"
go build
popd

pushd "$DSL_PATH/test"
cat src.t2flow | ../t2extract/t2extract result1
../t2compile/t2compile result1 > result2.t2flow
cat result2.t2flow | ../t2extract/t2extract result3
../t2compile/t2compile result3 > result4.t2flow

diff result2.t2flow result4.t2flow
diff result1 result3
popd
