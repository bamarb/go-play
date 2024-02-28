{{.}}
{{.Title}}
{{.HTML}}
{{.SafeHTML}}


<a title="{{.Title}}">
<a title="{{.HTML}}">

<a href="{{.HTML}}">
<a href="?q={{.HTML}}">
<a href="{{.Path}}">
<a href="?q={{.Path}}">

<!-- Encoding even works on non-string values! -->
<script>
  var dog = {{.Dog}};
  var map = {{.Map}};
  doWork({{.Title}});
</script>
