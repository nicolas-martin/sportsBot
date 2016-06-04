'use strict';
let exec = require('child_process').exec;

exports.handler = (event, context) => {
    const child = exec("./go_compiled_code "+event.firstname+" "+event.lastname+" "+event.phonenumber, (error) => {
        if(error){
            console.log('Error:',error);
            context.fail(error);
        }
    });
    child.stdout.on('data',function(data){
        console.log('Data:',data);
        context.succeed(data)
    });
};
