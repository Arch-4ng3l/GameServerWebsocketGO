<script>
    let url = "/api/logs"
    fetch(url, {
        method: "GET"
    }) 
    .then(response => {
        if (response.ok) {
            return response.json;
        }
    })
    .then(data => {
        console.log(data); 
    });
</script>



<div>
    <h2>Logs</h2>  
</div>


<style>

</style>
