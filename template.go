package main

var T0 = []byte(`<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		
		<title>`)

var T1 = []byte(`</title>
		
		<style type="text/css">
		body {
			padding: 0;
			margin: 0;
			font-family: "Helvetica Neue", Helvetica, sans-serif;
			color: #555;
		}
		.cont {
			width: 90%;
			max-width: 600px;
			margin: auto;
			margin-top: 20px;
		}
			.cont__label {
				font-size: 18px;
				margin-bottom: 5px;
				color: #999;
			}
			.cont__title {
				margin: 0;
				font-size: 32px;
				font-weight: bold;
				margin-bottom: 5px;
			}
			.cont__remaining {
				font-size: 22px;
			}
			.cont__description {
				line-height: 1.5;
				font-size: 18px;
			}
		</style>
	</head>
	<body>
		<div class="cont">
			<div class="cont__label">Coming Soon</div>
			<h1 class="cont__title">`)

var T2 = []byte(`</h1>
			<div class="cont__remaining">`)

var T3 = []byte(`</div>
			<div class="cont__description">`)

var T4 = []byte(`</div>
		</div>
	</body>
</html>`)
