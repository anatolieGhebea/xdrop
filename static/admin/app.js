(function() {

    let ip = document.querySelector('meta[name="ServerIp"]').content
    let port = document.querySelector('meta[name="ServerPort"]').content

    if( ip && port ){
        generateQRCode(`${ip}${port}`)
    } else {
        console.error('Server Ip or Port not provided');
    }
    

})()

function generateQRCode(url, schema = 'http'){
    console.log('qr');
    // from https://davidshimjs.github.io/qrcodejs/ 
    // dom might be slow, wait a few milliseconds before updating...
    const wait = setTimeout( () => {
        new QRCode("qrcode", {
            text: `${schema}://${url}`,
            width: 128,
            height: 128
        });
        clearTimeout(wait);
    }, 200 );
}