<script>
    import { writable } from "svelte/store";

    const url = "/api/conns"
    let conns = writable([])
    function getConnections() {
        fetch(url, {
            method: "GET",
        })
        .then(response => {
            if(response.ok) {
                return response.json();
            }
        })
        .then(data => {
            console.log(data);
            data.forEach(element => {
                conns.update(arr => [...arr, element]);
            });
            console.log(data[0].remote)
        })
        .catch(e => {
            console.log(e);
            return;
        });
    }
</script>

<div>
    <button on:click={getConnections} class="btn" >Get Connections</button>

    {#each $conns as {remote, local}}
        <div class="obj">
        <h2>remote: {remote}</h2>
        <h2>local : {local}</h2>
        </div>
    {/each}
</div>

<style>
.btn {
    border-style: solid;
    border-radius: 10px;
}
.obj{
    padding: 10px;
    outline: 3px;
    border-radius: 10px;
    border-style:solid;
    border-color:black;
    margin: 10px;
}

</style>
