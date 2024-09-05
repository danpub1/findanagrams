# gawk -M -f makedict.awk scowl-2020.12.07/final/*.* >allwords.go
function hash(s,   ii, rv, ss)
{
    rv = 1
    for (ii = 1; ii <= length(s); ii++)
    {
        ss = substr(s,ii,1)
        if (ss in prime)
        {
            rv *= prime[ss]
        }
    }

    return rv
}

BEGIN {
    prime["a"] = 2;
    prime["b"] = 3;
    prime["c"] = 5;
    prime["d"] = 7;
    prime["e"] = 11;
    prime["f"] = 13;
    prime["g"] = 17;
    prime["h"] = 19;
    prime["i"] = 23;
    prime["j"] = 29;
    prime["k"] = 31;
    prime["l"] = 37;
    prime["m"] = 41;
    prime["n"] = 43;
    prime["o"] = 47;
    prime["p"] = 53;
    prime["q"] = 59;
    prime["r"] = 61;
    prime["s"] = 67;
    prime["t"] = 71;
    prime["u"] = 73;
    prime["v"] = 79;
    prime["w"] = 83;
    prime["x"] = 89;
    prime["y"] = 97;
    prime["z"] = 101;
    prime["A"] = 2;
    prime["B"] = 3;
    prime["C"] = 5;
    prime["D"] = 7;
    prime["E"] = 11;
    prime["F"] = 13;
    prime["G"] = 17;
    prime["H"] = 19;
    prime["I"] = 23;
    prime["J"] = 29;
    prime["K"] = 31;
    prime["L"] = 37;
    prime["M"] = 41;
    prime["N"] = 43;
    prime["O"] = 47;
    prime["P"] = 53;
    prime["Q"] = 59;
    prime["R"] = 61;
    prime["S"] = 67;
    prime["T"] = 71;
    prime["U"] = 73;
    prime["V"] = 79;
    prime["W"] = 83;
    prime["X"] = 89;
    prime["Y"] = 97;
    prime["Z"] = 101;

    shortWords["a"] = 1
    shortWords["A"] = 1
    shortWords["i"] = 1
    shortWords["I"] = 1
    shortWords["o"] = 1
    shortWords["O"] = 1

    split("", found)

    localeMap["american"] = 16
    localeMap["australian_variant_1"] = 64
    localeMap["australian_variant_2"] = 128
    localeMap["australian"] = 32
    localeMap["british_variant_1"] = 512
    localeMap["british_variant_2"] = 1024
    localeMap["british_z"] = 2048
    localeMap["british"] = 256
    localeMap["canadian_variant_1"] = 8192
    localeMap["canadian_variant_2"] = 16384
    localeMap["canadian"] = 4096
    localeMap["english"] = 1
    localeMap["variant_1"] = 2
    localeMap["variant_2"] = 4
    localeMap["variant_3"] = 8

    partMap["abbreviations"] = 8
    partMap["proper-names"] = 16
    partMap["upper"] = 2
    partMap["words"] = 1
    partMap["contractions"] = 4

    sizeMap["10"] = 1
    sizeMap["20"] = 2
    sizeMap["35"] = 4
    sizeMap["40"] = 8
    sizeMap["50"] = 16
    sizeMap["55"] = 32
    sizeMap["60"] = 64
    sizeMap["70"] = 128
    sizeMap["80"] = 256
    sizeMap["95"] = 512

    print "package main"
    print ""
    print "// W is the type of the word data"
    print "type W struct {"
    print "    entry string"
    print "    dictSize uint16"
    print "    locale uint16"
    print "    part uint8"
    print "    low uint64"
    print "    high uint64"
    print "}"
    print ""
    print "var allWords = [...]W{"
}

function fmt(v)
{
    vX = sprintf("0x%x", v)
    vD = sprintf("%d", v)
    return length(vX) <= length(vD) ? vX : vD;
}

END {
    for (word in found) {
       hh = hash(word)
       if (hh / (0xFFFFFFFFFFFFFFFF+1) <= 0xFFFFFFFFFFFFFFFF)
       {
            fff = found[word]
            locale = fff % 32768
            fff = (fff - locale) / 32768
            part = fff % 32
            fff = (fff - part) / 32
            dictSize = fff

            printf("W{\"%s\",%s,%s,%s,%s,%s},", word, fmt(dictSize), fmt(locale), fmt(part), fmt(hh % (0xFFFFFFFFFFFFFFFF+1)), fmt(hh / (0xFFFFFFFFFFFFFFFF+1)));
            #printf("W{\"%s\",%s,%s,%s},", word, fmt(found[word]), fmt(hh % (0xFFFFFFFFFFFFFFFF+1)), fmt(hh / (0xFFFFFFFFFFFFFFFF+1)));
       }
    }

    print "}"
}


BEGINFILE {
    split(FILENAME, ff, ".")
    ext = sizeMap[ff[length(ff)]] * 32768 * 32
    split(ff[length(ff)-1], pp, "/")
    split(pp[length(pp)], nn, "-")
    if (length(nn) == 2 || length(nn) == 3)
    {
        if (length(nn) == 3) nn[2] = nn[2] "-" nn[3]
        locale = nn[1] in localeMap ? localeMap[nn[1]] : -1
        part = nn[2] in partMap ? partMap[nn[2]] * 32768 : -1
    }
    else
    {
        locale = -1
        part = -1
    }
} 

/^[a-zA-Z']+$/ {
    if (locale >= 0 && part >= 0)
    {
       if ((length($0) > 1 || $0 in shortWords) && $0 !~ /^[a-z]'[s]$/ && ($0 ~ /[aeiouyAEIOUY]/ || $0 ~ /^[A-Z']$/)) {
           if ($0 ~ /^[A-Z][a-z']+$/) $0 = tolower($0)
           if ($0 in found) found[$0] = or(found[$0], (locale + part + ext))
           else found[$0] = locale + part + ext
       }
    }
}