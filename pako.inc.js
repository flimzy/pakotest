if ( typeof($global.pako) === 'undefined' ) {
    try {
        $global.pako = require('pako');
    } catch(e) {
        throw("Cannot find global pako object. Did you load pako?");
    }
}
