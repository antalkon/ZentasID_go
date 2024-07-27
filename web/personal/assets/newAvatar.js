document.addEventListener("DOMContentLoaded", () => {
    const loadAvatarBtn = document.getElementById("wsbAva");
    const fileInput = document.getElementById("fileAvatar");

    loadAvatarBtn.addEventListener("click", () => {
        fileInput.click();
    });

    fileInput.addEventListener("change", async () => {
        const fileAvatar = fileInput.files[0];
        if (fileAvatar) {
            const formData = new FormData();
            formData.append('avatar', fileAvatar);

            try {
                const response = await fetch('/data/api/avatar', {
                    method: 'PUT',
                    body: formData
                });

                if (response.ok) {
                    console.log("File uploaded successfully");
                } else {
                    console.log("Error uploading file");
                }
            } catch (error) {
                console.error("Error uploading file:", error);
            }
        }
    });
});