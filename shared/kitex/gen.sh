
# Kitex-Gen
kitex -module tuzi-tiktok/kitex -I ..\..\idl\kitex ..\..\idl\kitex\comment.proto
kitex -module tuzi-tiktok/kitex -I ..\..\idl\kitex ..\..\idl\kitex\user.proto
kitex -module tuzi-tiktok/kitex -I ..\..\idl\kitex ..\..\idl\kitex\feed.proto
kitex -module tuzi-tiktok/kitex -I ..\..\idl\kitex ..\..\idl\kitex\message.proto
kitex -module tuzi-tiktok/kitex -I ..\..\idl\kitex ..\..\idl\kitex\publish.proto
kitex -module tuzi-tiktok/kitex -I ..\..\idl\kitex ..\..\idl\kitex\relation.proto
kitex -module tuzi-tiktok/kitex -I ..\..\idl\kitex ..\..\idl\kitex\favorite.proto

# Service 脚手架
kitex -module tuzi-tiktok/service/auth     -service auth-api     -use tuzi-tiktok/kitex/kitex_gen -I ..\..\..\idl\kitex ..\..\..\idl\kitex\user.proto
kitex -module tuzi-tiktok/service/comment  -service comment-api  -use tuzi-tiktok/kitex/kitex_gen -I ..\..\..\idl\kitex ..\..\..\idl\kitex\comment.proto
kitex -module tuzi-tiktok/service/favorite -service favorite-api -use tuzi-tiktok/kitex/kitex_gen -I ..\..\..\idl\kitex ..\..\..\idl\kitex\favorite.proto
kitex -module tuzi-tiktok/service/feed     -service feed-api     -use tuzi-tiktok/kitex/kitex_gen -I ..\..\..\idl\kitex ..\..\..\idl\kitex\feed.proto
kitex -module tuzi-tiktok/service/message  -service message-api  -use tuzi-tiktok/kitex/kitex_gen -I ..\..\..\idl\kitex ..\..\..\idl\kitex\message.proto
kitex -module tuzi-tiktok/service/publish  -service publish-api  -use tuzi-tiktok/kitex/kitex_gen -I ..\..\..\idl\kitex ..\..\..\idl\kitex\publish.proto
kitex -module tuzi-tiktok/service/relation -service relation-api -use tuzi-tiktok/kitex/kitex_gen -I ..\..\..\idl\kitex ..\..\..\idl\kitex\relation.proto

