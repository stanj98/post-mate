
async function createNote() {
    var title = document.querySelector("#input-title").value;
    var trix = document.querySelector(".trix-custom-content");
    console.log(trix);
    const editorContent = trix.editor.getDocument().toString();
    console.log(trix.editor.getDocument());
    const tempDiv = document.createElement('div');
    tempDiv.innerHTML = editorContent;
    const images = tempDiv.querySelectorAll('img');
    console.log(images);
    const imageUrls = [];
    images.forEach((img) => {
        imageUrls.push(img.src);
    });
    console.log(imageUrls);
    //need to extract data from trix content and store
    //need to extract attachments and save via AWS s3 bucket
    const data = {
        title: title,
        body: "<div> Dummy content </div>",
    }
    const headers = {
        'Content-Type': 'application/json'
    }
    await axios.post(ROOT_API + "api/notes/", data, headers)
    .then(response => {
        console.log(response);
        if (response.status == 201) {
            window.location.href = ROOT_API + "view-notes";
        }
    });
}