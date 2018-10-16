// part 1 of multiply strings
// https://www.codewars.com/kata/55911ef14065454c75000062
function multiply_v1(a, b) {
    let ar = a.replace(/^0+/, "").split("").reverse();
    let br = b.replace(/^0+/, "").split("").reverse();
    let r = [];
    for (let ai in ar) {
        let av = ar[ai];
        for (let bi in br) {
            let bv = br[bi];
            let i = parseInt(ai) + parseInt(bi);
            let nexti = i + 1;
            r[i] = r[i] ? r[i] : 0;
            r[i] += av * bv;
            if (r[i] > 9) {
                r[nexti] = r[nexti] ? r[nexti] : 0;
                r[nexti] += Math.floor(r[i] / 10);
                r[i] = r[i] % 10
            }
        }
    }
    return r.reverse().join("");
}

//https://www.codewars.com/kata/multiplying-numbers-as-strings-part-ii
function multiply(n, o) {
    //process negative
    let negative = ((n[0] === "-") && (o[0] !== "-")) || ((n[0] !== "-") && (o[0] === "-"));
    n = n.replace("-", "").replace(/^0+/, "");//process zero like 00.0020
    o = o.replace("-", "").replace(/^0+/, "");
    //process decimals
    let fix = 0;
    if (n.indexOf(".") >= 0) {
        n = n.replace(/0+$/, ""); //process 0.0200
        fix += n.length - n.indexOf(".") - 1;
        n = n.replace(".", "");
    }
    if (o.indexOf(".") >= 0) {
        o = o.replace(/0+$/, "");
        fix += o.length - o.indexOf(".") - 1;
        o = o.replace(".", "");
    }
    if (n === "" || o === "") {
        return "0";
    }
    let r = multiply_v1(n, o);

    r = Array(fix > r.length ? fix - r.length + 1 : 0).fill(0).join("") + r;

    if (r.length >= fix && fix > 0) {
        r = r.substr(0, r.length - fix) + "." + r.substr(r.length - fix, fix)
    }

    r = negative ? "-" + r : r;
    r = r.indexOf(".") >= 0 ? r.replace(/0+$/, "").replace(/\.$/, "") : r;

    return r
}

function doV1Test(a, b, expect) {
    r = multiply_v1(a, b);
    if (r !== expect) {
        console.log("FAILED expect [" + r + "] to equal [" + expect + "]");
        return
    }
    console.log("SUCCESS [" + a + "*" + b + "," + expect + "]");
}

function doTest(a, b, expect) {
    r = multiply(a, b);
    if (r !== expect) {
        console.log("FAILED expect [" + r + "] to equal [" + expect + "]");
        return
    }
    console.log("SUCCESS [" + a + "*" + b + "," + expect + "]");
}

// doTest("00123", "2", "246");
doTest("2.5", "4", "10");
doTest("0.0008", "0.01000000", "0.000008");
doTest("-0.01", "0.0000", "0");
doTest("30", "69", "2070");
doTest("-5.0908", "-123.1", "626.67748");
doTest("-135", "26", "-3510");
doTest("0.05", "12345", "617.25");
doTest("2.01", "3.0000", "6.03");
doTest("0.05", "0.12345", "0.0061725");
doTest("0.050", "00.1235000", "0.006175");
// doTest("-377558706748806277773288244838373888702242753209392770149690306008920756023582179370956756559777.25559856343932958556193953881966509092244759243150567992368341024439630786634467694925457698990265",
//     "-8608581268198563355162171931743415486197120972410363902236237182415833124768008542562040220861229.45017778088907084489472186529363266145797555456888475612",
//     "3250244810563048227782667476859948824696649320164283472191455935028021272211676440888518019069294002811313461792891242778245080565931066469881408183403359947323071880995048624300848651420027722.822894980542679642930666219118102689156779986042644343992950454764503836767667739387629571439672101693688405372940433162376559136574844186512632747791718");
// doV1Test("1020303004875647366210", "2774537626200857473632627613", "2830869077153280552556547081187254342445169156730");
// doV1Test("58608473622772837728372827", "7586374672263726736374", "444625839871840560024489175424316205566214109298");