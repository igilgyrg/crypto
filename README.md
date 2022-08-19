# rest api crypto

# LTC service

GET /candles - list of candles -- 200, 404, 500 \
POST /candles - create new candle --204, 400, 500 \
GET /users/?datetime - get candle by datetime -- 200, 404, 500 \
GET /users/?start&end - list of candles between time period -- 200, 404, 500 \
