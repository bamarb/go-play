{{ define "header" }}
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8" />
        <meta
            name="viewport"
            content="width=device-width, initial-scale=1"
        />
    </head>
    <body>
      <h1> Snippet Rendered By Header </h1>
{{end}}

{{ define "footer" }}
  <div>
   <p> Made in {{.country}} </p>
  </div>
 </body>
 </html>
{{end}}
