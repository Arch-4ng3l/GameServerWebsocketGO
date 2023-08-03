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
            conns.set([]);
            data.forEach(element => {
                conns.update(arr => [...arr, element]);
            });
        })
        .catch(_ => {
            return;
        });
    }

    function closeConn(remoteAddr, localAddr) {
        if (remoteAddr) {
            let obj = {
                remote: remoteAddr,
                local: localAddr,
            };
            let jsonObj = JSON.stringify(obj);
            fetch(url, {
                method: "POST",
                body: jsonObj,
            })
            .then(response => {
                if(response.ok) {
                    console.log("removed");

                    conns.update(arr => arr.filter(obj => obj.remote !== remoteAddr && obj.local !== localAddr));
                }
            });

        }
    }
</script>

<div>
    <button on:click={getConnections} class="btn" >Get Connections</button>
    <div class="list">
    {#each $conns as {remote, local}}
        <div class="obj">
        <div class="text">
        <span id="remote" ><b>remote: </b></span>
        <span>&emsp;</span>
        <span class="addr"><b>{remote}</b></span>
        <br>
        <span id="local" ><b>local : </b></span>
        <span>&emsp;</span>
        <span class="addr"><b>{local}</b></span>
        <br>
        </div>
        <br>
        <button on:click={() => closeConn(remote, local)} class="closebtn" >Close Connection</button>
        </div>
    {/each}
    </div>
</div>

<style>

.text {
    text-align: left;
}
.addr {
    float: right;
    text-align: right;
}
.closebtn {
    border-color: red;
    border-radius: 10px;
    transition: 0.4s;
}
.closebtn:hover {
    transition: 0.4s;
    color: red;
    border-width: 3px;
    font-weight: bold;
}

.btn {
    border-style: solid;
    border-radius: 10px;
}
.btn:hover {
    border-color: #ff3e00;
    font-weight: bold;
    color: #ff3e00;
}

    
.list {
    margin-left: 30%;
    width: 40%;
    padding: 10px;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    right: 1em;
    align-content: center;
}

.obj{
    padding: 10px;
    outline: 3px;
    border-radius: 10px;
    border-style:solid;
    border-color:black;
    margin: 10px;
}
.obj:hover{
    border-color: #ff3e00;
}

</style>
