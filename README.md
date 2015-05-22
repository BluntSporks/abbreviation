# abbreviation
Reusable Golang library to provide functions to check the validity of abbreviations

## Purpose
To list all abbreviations that I use in programs for reference.

## Status
Ready to use

## Rules
### Sound
* Abbreviations should be pronounceable.
* Abbreviations should have at least one vowel.
* Abbreviations should not split up plosive/liquid combinations, for example, the pr in represent or the tl in template.
* Abbreviations should split up a plosive/plosive combination, for example, the ct in dictionary or pt in caption,
  especially where the original t is pronounced sh in tion.

### Spelling
* Abbreviations should not have more than three consonants in a row.
* Abbreviations should usually end in a consonant, unless the vowel is needed for discrimination, for example, alg and algo.
* All of the letters in the abbreviation should be present in the long form and in the same order.
* The letters of the abbreviation need not appear in sequence in the long form, for example, recv and receive.

### Length
* A abbreviation should be less than or equal to half the length of the original form.
* Abbreviations should be at least three letters long.

### Meaning
* Abbreviations should not be whole words that mean something else.
* Abbreviations should not just consist of the prefix of a word, for example, sym for symbol or syl for syllable.

### Interpretation
* Abbreviations shouldn't be ambiguous. However, if the names are different that no confusion can result, they are OK.

### Origin
* Abbreviations should be compatible with the Google Go language. They should be ideally already used in Go code.
* Abbreviations should be used for common words, since it little helps to abbreviate rare words.
* Abbreviations should be verified at [AllAcronyms.com](http://www.allacronyms.com/) as being in common use.
* Abbreviations should come from the world of IT, technology, and computers, and already be used there.

### Endings
It is possible to use a list of suffix abbreviations to discriminate between different forms of a word if you need to.
However, it is not necessary to do so in every case, only in cases where you want to discriminate between two different
forms in the same context, for example, to use a plural form to identify an array and a singular form to identify an
element of the same array (the ansz array could contain ans elements). Or you might care about the difference between
emple, emplr, and emplt in the same application or database.

The following list is recommended:
* b for able
* c for ence or ance
* d for ed
* e for ee
* g for ing
* n for tion or sion
* l for ial
* r for er or or
* s for plural of abbreviation not ending in s
* t for ment
* v for ive
* z for plural of abbreviation ending in s

Examples include:
* emplb = employable
* distc = distance
* empld = employed
* emple = employee
* emplg = employing
* authn = authentication or authorization
* binl = binomial
* emplr = employer
* emplrs = employers
* emplt = employment
* ansz = answers 

### Acronyms
* Acronyms are OK in identifier names, but the acronym should already be well known. They are not listed here.
* Acronyms should be in consistent case, all-lower or all-upper, as recommended in
  [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments). I prefer all caps for all acronyms,
  for example, ServeHTTP or HTTPServer.
* Do not run acronyms together, for example, SQLID or FTPHTTP, since it is hard to read.
* Acronyms should be at least three letters long.

### Exceptions
There are a few exceptions to the above rules for common, well-established forms.
* ct and pt can be used for ction and ption if the abbreviation would be too short otherwise, for example, act and opt.
* x can be used for trans even though it is not a letter of the word.
* Some whole words are OK, such as fact for factorial.

## Limitations
This style of abbreviation is intended for use in naming classes, functions and variables. However there are other
types of prefixing, for example the three-letter prefixes used to distinguish field names in the same database table.
Examples would include cusID for customer ID and ordID for order ID. Those prefixes need not follow the same rules as
these and are not covered here.

## List
* abbrev = abbreviate, abbreviation
* abs = absolute
* acc = accuracy, accurate, account
* accum = accumulate, accumulation, accumulator
* act = actual
* addr = address
* aggr = aggregate, aggregation
* alg = algebra
* algo = algorithm
* alloc = allocate, allocation
* amp = ampersand
* anc = ancestor
* ans = answer
* apos = apostrophe
* approx = approximate, approximation
* arch = architect, architecture
* arg = argument
* asc = ascending
* assoc = associative, associate
* attr = attribute
* auth = authenticate, authentication, authorize, authorization
* avg = average
* bin = binary, binomial
* biz = business
* bool = boolean
* buf = buffer
* cal = calendar
* calc = calculate, calculation, calculator, calculus
* cap = capitalize, capitalization, capacity
* ceil = ceiling
* cert = certify, certificate, certification
* chan = channel
* char = character
* coef = coefficient
* col = column
* coll = collect, collection
* comb = combine, combination
* comp = compare, comparison, compensate, compensation
* cond = condition
* config = configure, configuration
* conn = connect, connection, connector
* cont = continue, continuation, contain, container
* conv = convert, conversion
* cop = complex
* corr = correlate, correlation
* covar = covariance
* crypt = cryptography, cryptology, cryptanalysis, cryptosystem
* cur = current
* cust = customer
* dec = decimal, decrement, decode, decrypt
* decl = declare, declaration
* def = define, definition
* deg = degree
* del = delete, deletion
* delim = delimiter
* den = denomination, denominator
* dep = depend, dependent
* dept = department
* desc = descendant, descending, describe, description
* deser = deserialize, deserialization
* dest = destination
* det = determinant
* dev = deviate, deviation, develop, development
* diag = diagonal, diagnosis, diagnostic
* diam = diameter
* dic = dictionary
* diff = difference
* dig = digest
* dim = dimension
* dir = directory, directive, direct, direction
* disc = discriminant
* dist = distribute, distribution, distributor, distance
* div = divide
* doc = document, documentation
* dom = domain
* dupe = duplicate, duplication
* elem = element
* empl = employee, employment
* enc = encode, encrypt
* enum = enumeration
* env = environment
* equiv = equivalent, equivalence
* err = error
* esc = escape
* est = estimate, estimation
* eval = evaluate, evaluation
* excl = exclude, exclusion, exclamation
* excp = except, exception
* exec = execute, execution
* exp = exponent, expect
* expr = express, expression
* ext = extend, extension, external
* fact = factorial
* fem = feminine
* fig = figure
* float = floating
* frac = fraction
* freq = frequent, frequency
* func = function
* gen = generate, generation, generator, general, generic
* geo = geography, geographic
* geom = geometric
* guar = guarantee
* har = harmonic
* hex = hexadecimal
* horiz = horizontal
* ident = identification, identifier
* imag = imaginary
* imm = immediate
* impl = implement, implementation
* inc = increment
* incl = include, inclusion
* ind = indirect, indicate, indication
* indep = independent
* inf = infinite, infinity
* info = information
* init = initial, initialize, initialization
* ins = insert, insertion
* inst = instance, instantiate, instruction
* int = integer, internal
* intc = intercept
* intf = interface
* inv = inverse
* iter = iterate, iteration, iterator
* kurt = kurtosis
* lang = language
* len = length
* lex = lexicon
* lib = library
* lim = limit, limitation
* lin = linear
* lit = literal
* loc = locate, location, locale
* log = logarithm
* low = lowercase
* mant = mantissa
* masc = masculine
* max = maximum
* med = medium, median
* mem = memory
* min = minimum, minute
* mod = modulo, modify, modification, module
* mul = multiply
* multi = multiple
* mut = mutate, mutation, mutable
* mutex = mutual exclusion
* nat = natural
* neg = negate, negation, negative
* net = network
* num = number, numeration, numerator
* obj = object
* oct = octal
* op = operate, operation, operator
* opt = option
* orig = original, originate, origination
* ovfl = overflow
* pal = palette
* par = parent
* para = paragraph
* param = parameter
* pat = pattern
* perf = perform, performance
* perim = perimeter
* perm = permutation, permanent
* perp = perpendicular
* phon = phonetic
* pic = picture
* pos = positive, position
* poss = possible, possibly
* prec = precise, precision
* pref = prefer, preference
* prep = prepare, preparation
* prev = previous
* pri = private, primary
* prio = priority, prioritize
* prob = probability
* proc = procedure, process
* prod = product
* prof = profile
* prog = program, progress
* proj = project
* pron = pronounce, pronunciation
* prop = proportion
* prot = protocol
* proto = prototype
* pub = public
* punc = punctuate, punctuation
* quad = quadratic
* ques = question
* quot = quotient
* rad = radian, radius, radical
* rand = random
* rec = recursive, record
* rect = rectangle
* recv = receive
* ref = reference
* reg = register, regular
* regexp = regular expression
* rel = relate, relation, relative
* rem = remainder, remove
* ren = rename
* rep = repeat, repetition
* repl = replace, replacement
* repo = repository
* repr = represent, representation, representative
* req = require, requirement
* res = resource, result, resolve, resolution
* resp = respond, response, responsive
* ret = return
* rev = reverse
* san = sanitize, sanitization
* sched = schedule
* sci = science, scientific
* sec = second, secondary, secure, security, section
* seg = segment
* sel = select, selection
* sem = semaphore
* sent = sentence
* sep = separate, separation, separator
* seq = sequence
* ser = serial, serialize, serialization, series
* sig = signal, signature
* sim = simulate, simulation
* sing = singular
* sock = socket
* spec = specify, specification, specific
* stat = statistic
* sub = subtract
* syll = syllable
* symm = symmetry
* sync = synchronized
* sys = system
* tan = tangent
* tech = technology
* temp = temporary
* term = terminate, termination, terminal
* thru = through
* tok = token, tokenize
* trig = trigonometry
* up = uppercase
* upd = update
* util = utility
* val = value, validate, validation
* var = variance, variable
* vec = vector
* ver = version, verify, verification
* vert = vertical
* vis = visible, visibility
* vol = volume
* win = window
* xfer = transfer
* xlat = translate, translation
* xmit = transmit
* xor = exclusive or
* xpar = transparent, transparency
* xpos = transpose, transposition

## Not in combination
These are OK to use as whole words, or at the beginning or end of a camelcased word, but try not to use them in the
middle of words due to their lack of vowels.
* cmd = command
* cmp = compare, comparison
* cnt = count
* ctl = control
* ctx = context
* cvt = convert
* dst = destination
* fld = field
* fmt = format
* fwd = forward
* hdl = handle
* hdr = header
* i18n = internationalization
* l10n = localization
* msg = message
* pkg = package
* ptr = pointer
* qry = query
* src = source
* srv = server
* std = standard
* stm = statement
* str = string
* tbl = table

## Not abbreviated
Here is a list of words that I do not abbreviate, even though some of them already have common abbreviations like line.
* array
* default
* empty
* equal
* file
* index
* line
* name
* test
* text
