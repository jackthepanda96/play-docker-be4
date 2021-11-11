# Belajar Docker
Repository belajar docker ALTA Immerseive Back-End Batch 4

Untuk materi ini teman-teman bisa download docker sesuai dengan OS masing-masing:

* Windows [link](https://docs.docker.com/desktop/windows/install/)
* Mac [link](https://docs.docker.com/desktop/mac/install/)
* Linux 
    ```sh
        sudo apt-get install docker.io
    ```

Selain installasi docker teman-teman juga harus membuat account pada:

* [DockerHub](https://hub.docker.com/)
* [Amazone Web Service](https://aws.amazon.com/)
    * Untuk melakukan pendaftaran memerlukan rekening Jenius atau Bank Jago.

## Docker Command
<table>
    <thead>
        <tr>
            <td> Command</td>
            <td> Fungsi</td>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td> docker build `path/to/Dockerfile`</td>
            <td> Create image dari Dockerfile yang telah dibuat</td>
        </tr>
        <tr>
            <td> docker pull `image_name`</td>
            <td> Download images yang tersedia pada docker hub</td>
        </tr>
        <tr>
            <td> docker images</td>
            <td> Show seluruh images yang telah didownload</td>
        </tr>
        <tr>
            <td> docker rmi `image_id`</td>
            <td> Delete image. (Dalam beberapa kasus, images tidak bisa langsung dihapus. Gunakan tambahan `-f`</td>
        </tr>
        <tr>
            <td> docker run `image_name`</td>
            <td> Untuk melakukan *CREATE* create container sekaligus *START* container dari image yang dipilih</td>
        </tr>
        <tr>
            <td> docker ps</td>
            <td> Show seluruh container</td>
        </tr>
        <tr>
            <td> docker start `container_name/container_id`</td>
            <td> Untuk *START* container</td>
        </tr>
        <tr>
            <td> docker stop `container_name/container_id`</td>
            <td> Untuk *STOP* container</td>
        </tr>
        <tr>
            <td> docker rm `container_name/container_id`</td>
            <td> Delete container. Biasakan untuk STOP container baru lakukan delete</td>
        </tr>
        <tr>
            <td> docker logs `container_name/container_id`</td>
            <td> Untuk menampilkan logs yang terdapat pada container</td>
        </tr>
    </tbody>
</table>
