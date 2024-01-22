<?php

class Serializer
{
    public function serializerNumber(int $number)
    {
        return serialize($number);
    }

    public function serializerString(string $str)
    {
        return serialize($str);
    }
}
