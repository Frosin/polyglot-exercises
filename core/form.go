package core

var Form = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/bootstrap.min.css"
        integrity="sha384-TX8t27EcRE3e/ihU7zmQxVncDAy5uIKz4rEkgIXeMed4M0jlfIDPvg6uqKI2xXr2" crossorigin="anonymous">
    <title>Полиглот-упражнения</title>
</head>

<body>
    <div class="container">
        <h4>Введите перевод фразы:</h4>
        <form>
            <div class="form-group">
                <label for="formGroupExampleInput">%phrase%</label>
				<input type="text" class="form-control is-invalid" name="translate" placeholder="Введите перевод">
				<div class="invalid-feedback">%correct%</div>
				<input type="hidden" class="form-control" name="phrase" value="%phrase%">
				<input type="hidden" class="form-control" name="result" value="%result%">
            </div>
            <button type="submit" class="btn btn-primary">Проверить</button>
        </form>
    </div>
</body>

</html>
`
